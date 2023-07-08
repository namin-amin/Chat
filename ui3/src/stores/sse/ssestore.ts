import {create} from "zustand";
import {SSE} from "../../models/sse";
import {NET_URL} from "../../constants/netwrok";
import {useAuthStore} from "../auth/authstore";

interface Store {
    sseClient?: SSE,
}

interface Action {
    init: () => void;
    dispose: () => void;
}

export const useSSEStore = create<Store & Action>()((set, get) => ({
    sseClient: undefined,
    init: () => {
        if (get().sseClient === undefined) {
            set({
                sseClient: new SSE(NET_URL + "/chats/sse", useAuthStore.getState().userid)
            })
        }
    },
    dispose: () => {
        get().sseClient?.close();
    }
}));
