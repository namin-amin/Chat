import {AddConversation} from "../components/AddConversation";
import {ChatList} from "./ChatList.tsx";
import SideBar from "../components/SideBar";
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {Settings} from "./Settings";
import {ChatWindow} from "../components/ChatWindow.tsx";

const routes = createBrowserRouter([
    {
        path: "/ui",
        element: <SideBar/>,
        children: [
            {
                path: "chats",
                element: <ChatList/>,
                children: [
                    {
                        path: ":id",
                        index: true,
                        element: <ChatWindow/>,
                    },
                ],
            },
            {
                path: "newchat",
                element: <AddConversation/>,
            },
            {
                path: "settings",
                element: <Settings/>,
            },
        ],
    },
]);

export const Layout = () => {
    return <RouterProvider router={routes}/>;
};
