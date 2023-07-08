import {create} from "zustand";
import {Chat} from "../../models/conversation";
import {client} from "../../utils/client";
import {useConversationStore} from "./conversationstore";
import {NewChatDto} from "../../models/messages.ts";

interface State {
    currentChats: Chat[]
}

interface Action {
    addNewChat: (chat: Chat) => void
    getChatsWithPagination: (conversationId: string, pageNumber: string) => void;
    setCurrentChats: (chats: Chat[]) => void
    sendNewChat: (newChat: NewChatDto) => Promise<void>
}

export const useChatStore = create<State & Action>()((set, get) => ({
    currentChats: [],
    addNewChat: (chat) => {
        if (chat.conversationId === useConversationStore.getState().currentConversationId) {
            set((state) => ({
                currentChats: [...state.currentChats, chat]
            }))
            useConversationStore.getState().addChatsToConversation(get().currentChats, chat.conversationId, true)
        }
    },
    //TODO need to fetch conversations when req
    getChatsWithPagination: async (conversationId: string, pageNumber: string) => {
        const result = await chatsWithPagination(conversationId, pageNumber)
        if (result !== undefined && result.length > 0) {
            const dd: Chat[] = []

            //TODO way inefficient fix the pagination in server
            result.forEach(cht => {
                let toAdd = false
                get().currentChats.forEach(curCht => {
                    if (curCht.id === cht.id) {
                        toAdd = true
                    }
                })
                if (!toAdd) {
                    dd.push(cht)
                }
            })

            if (dd.length > 0) {
                set(() => ({
                    currentChats: [...get().currentChats, ...dd]
                }))
                useConversationStore.getState().addChatsToConversation(get().currentChats, conversationId, true)
            }
        }
    },
    setCurrentChats: (chats) => {
        set({
            currentChats: chats
        })
    },
    sendNewChat: async (newChat) => {
        const response = await client.post("/chats/message", newChat);
        if (response.status === 202) {
            const addedChat: Chat = response.data;
            if (addedChat !== undefined) {
                get().addNewChat(addedChat)
            }
        }
    }
}))

async function chatsWithPagination(conversationId: string, pageNumber: string) {
    conversationId = conversationId === "" ? "20" : conversationId
    pageNumber = conversationId === "" ? "30000" : pageNumber
    const response = await client.get(`/chats/paginate/?pagecount=${pageNumber}&itemsperpage=${100}&convid=${conversationId}`)
    return response.data as Chat[]
}