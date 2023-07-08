import {FC} from "react";
import {User} from "../models/user";
import {Avatar} from "./Avatar";
import {Button} from "./Button";
import {BsPlus} from "react-icons/bs";
import {useConversationStore} from "../stores/chat/conversationstore";
import {useNavigate} from "react-router-dom";

interface AddConversationProps {
    user: User;
}

export const Person: FC<AddConversationProps> = ({user}) => {
    const makeNewConversation = useConversationStore(
        (state) => state.addConversation
    );

    const navigator = useNavigate();

    const addNewConversationAndNavigate = async () => {
        const conversation = await makeNewConversation(user.id);
        console.log("from person" + conversation);

        if (conversation !== undefined) {
            console.log(conversation);

            navigator(`/ui/chats/${conversation.receiverId}`); //todo error if initiator is other use
        }
    };

    return (
        <div
            className="
            box-border
            w-80
            bg-darksecondary
            rounded-md
            flex flex-col
            items-center
            h-60
            p-2 gap-2
            text-neutral-100
            transition-all
            hover:scale-105">
            <Avatar AvatarName={user.name} className="items-start w-28 h-28"/>
            <h1 className="text-2xl">{user.name}</h1>
            <Button
                onClickFun={addNewConversationAndNavigate}
                text="NewChat"
                className="bg-tirtiarry rounded-md text-lg"
                icon={BsPlus}
            />
        </div>
    );
};
