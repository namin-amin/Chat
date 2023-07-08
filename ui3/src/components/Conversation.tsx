import {FC} from "react";
import {twJoin} from "tailwind-merge";
import {Avatar} from "./Avatar";
import {Link} from "react-router-dom";

interface ConversationProps {
    className?: string;
    avatarImage?: string;
    name: string;
    senderId: string
    description?: string;
}

export const Conversation: FC<ConversationProps> = ({
                                                        name,
                                                        senderId,
                                                        avatarImage,
                                                        className,
                                                        description,
                                                    }) => {


    function AvatarSet(avatarImg: string | undefined, name: string) {
        if (avatarImg !== "" && avatarImg !== undefined) {
            return <img src={avatarImg} alt="avatar" className="flex-none"/>;
        }
        return <Avatar AvatarName={name} className="float-none"/>;
    }

    return (
        <Link
            className={twJoin([
                `text-neutral-400
                 w-full 
                 max-h-32 
                 flex 
                 items-center 
                 gap-2 p-2 m-2 
                 transition 
                 rounded-md 
                 hover:bg-bgmain 
                 cursor-pointer`,
                className,
            ])}
            to={senderId}
        >
            {AvatarSet(avatarImage, name)}
            <div className="h-full w-[70%] flex flex-col box-border flex-1 ">
                <h2 className="text-xl  text-white">{name}</h2>
                <p className=" flex-1 w-full truncate overflow-ellipsis ">
                    {description}
                </p>
            </div>
        </Link>
    );
};
