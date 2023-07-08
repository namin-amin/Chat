import {create} from 'zustand';
import {Chat, conversation} from '../../models/conversation';
import {useAuthStore} from '../auth/authstore';
import {client} from '../../utils/client';
import {useChatStore} from "./chatstore.ts";

interface State {
    conversations: conversation[];
    currentConversationId: string;
}

interface Action {
    init: () => Promise<void>;
    addConversation: (senderId: string) => Promise<conversation | undefined>;
    addChatsToConversation: (chats: Chat[], convId: string, completeReplace?: boolean) => void;
    populateCurrentChats: (id: string) => void;
}

export const useConversationStore = create<State & Action>()((set, get) => ({
    conversations: [],
    currentConversationId: "",
    addConversation: async (senderId: string) => {

        const newConv = await addNewConversation(senderId);
        if (newConv !== undefined) {
            const newData = [...get().conversations, newConv];
            set({
                conversations: newData
            });
        }

        return newConv;
    },
    init: async () => {
        if (get().conversations === undefined || get().conversations.length <= 0) {
            const data = await initConversations();
            console.log(data);

            set({
                conversations: data
            });
        }
    },
    addChatsToConversation: (chats, convId, completeReplace = false) => {
        const currentConversation = get().conversations.find((c) => {
            return c.id === convId
        })

        if (currentConversation === undefined) {
            return
        }
        const conversationAlreadyThere = get().conversations.indexOf(currentConversation)

        if (completeReplace) {
            currentConversation.Chats = [...chats]
        } else {
            currentConversation.Chats = [...currentConversation.Chats, ...chats]
        }

        if (conversationAlreadyThere === -1) {
            console.log("adding to new con")
            set(state => ({
                conversations: [...state.conversations, currentConversation]
            }))
        } else {
            console.log("trying to add conversation  again")
        }


    },
    populateCurrentChats: (id) => {
        const currentConversation = get().conversations.find(conversation => {
            return (conversation.senderId === id || conversation.receiverId === id)
        })
        if (currentConversation !== undefined && get().currentConversationId !== currentConversation.id) {
            useChatStore.getState().setCurrentChats(currentConversation.Chats)
            set({
                currentConversationId: currentConversation.id //check
            })
        }
    }
}));


const initConversations = async (): Promise<conversation[]> => {
    try {
        const response = await client.get(`/conversation/sender/${useAuthStore.getState().userid}`);
        if (response.status == 200) {
            return response.data satisfies conversation[];
        }
    } catch (error) {
        console.log("could not get the conversations" + error);
    }
    return [];
};


const addNewConversation = async (senderId: string): Promise<conversation | undefined> => {

    const response = await client.post("/conversation/", {
        senderId: useAuthStore.getState().userid,
        receiverId: senderId,
    });

    if (response.status === 202) {
        console.log(response.data);
        return response.data;
    }
};