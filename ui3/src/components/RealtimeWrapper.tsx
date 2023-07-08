import React, {FC, useEffect, useMemo} from "react"
import {useSSEStore} from "../stores/sse/ssestore"
import {MessageTypes} from "../models/messages.ts";
import {useChatStore} from "../stores/chat/chatstore.ts";


interface RealtimeWrapperProps {
    children: React.ReactNode
}

const RealtimeWrapper: FC<RealtimeWrapperProps> = ({
                                                       children
                                                   }) => {

    const [sseStore, initSSE] = useSSEStore(state => ([state.sseClient, state.init]));
    const addNewChat = useChatStore(state => state.addNewChat);

    useEffect(() => {
        initSSE()
    }, [initSSE])


    useMemo(() => {

        if (sseStore !== undefined) {
            sseStore.Subscribe(MessageTypes.DMessage, (ev: MessageEvent) => {
                const receivedData = JSON.parse(ev.data)
                console.log(JSON.parse(receivedData.data))
                addNewChat(JSON.parse(receivedData.data))
            })
        }

    }, [addNewChat, sseStore])


    return (
        <>
            {children}
        </>
    )
}

export default RealtimeWrapper