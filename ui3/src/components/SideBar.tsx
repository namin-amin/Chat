import React, {useEffect} from "react";
import SideBarItem from "./SideBarItem";
import {AiOutlineUser, AiOutlineUsergroupAdd} from "react-icons/ai";
import {BsChatLeft} from "react-icons/bs";
import {Outlet} from "react-router-dom";
import {useSSEStore} from "../stores/sse/ssestore";

const SideBar: React.FC = () => {
    const initSSE = useSSEStore((state) => state.init);

    useEffect(() => {
        initSSE();
    }, [initSSE]);

    return (
        <main
            className="
                    h-full
                    w-full
                    flex
                    flex-1
                    flex-row
                    bg-bgmain"
        >
            <div
                className="
                          flex
                          flex-col
                          h-full
                          w-14
                          bg-bgsecondary
                          rounded-md
                          p-1
                          justify-between
                          items-center"
            >
                <div className="
                                place-self-start
                                p-2
                                flex flex-col
                                w-full
                                mb-1
                                justify-center
                                items-center"
                >
                    <SideBarItem route="chats" name="Chats" icon={BsChatLeft}/>
                    <SideBarItem
                        route="newchat"
                        name="People"
                        icon={AiOutlineUsergroupAdd}
                    />
                </div>

                <div className="place-self-end flex flex-col w-full mb-1 justify-center items-center">
                    <SideBarItem route="settings" icon={AiOutlineUser}/>
                </div>
            </div>
            <section className="h-full w-full mx-2">
                <Outlet/>
            </section>
        </main>
    );
};

export default SideBar;
