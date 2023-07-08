export interface Message extends MessageData {
    senderId: string;
    receiverId: string;
    data: string;
}

export interface MessageData {
    event: MessageTypes;
    retry: number;
}

export enum MessageTypes {
    Open = "openConnection",
    Close = "close",
    DMessage = "message",
    Broadcast = "broadcast",
    NewConn = "newConnection"
}

export interface NewChatDto extends Message {
    conversationId: string;
}