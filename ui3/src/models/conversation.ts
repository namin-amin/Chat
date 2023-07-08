import {BaseModel} from "./base";
import {Message} from "./messages";
import {User} from "./user";

export interface conversation extends BaseModel {
    senderId: string;
    receiverId: string;
    Chats: Chat[];
    sender: User;
    receiver: User;

}

export interface Chat extends BaseModel, Message {
    conversationId: string;
}
