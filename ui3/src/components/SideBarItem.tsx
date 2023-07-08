import React from "react";
import {IconType} from "react-icons";
import {twMerge} from "tailwind-merge";
import {NavLink} from "react-router-dom";

interface SideBarItemProps {
    className?: string;
    children?: React.ReactNode;
    name?: string;
    route: string;
    icon: IconType;
}

const SideBarItem: React.FC<SideBarItemProps> = ({
                                                     name,
                                                     children,
                                                     className,
                                                     icon: Icon,
                                                     route,
                                                 }) => {
    return (
        <NavLink
            to={`${route}`}
            className={(isActive) => {
                return twMerge([
                    `flex flex-col 
                    w-full 
                    justify-center 
                    items-center m-2 
                    text-neutral-500 
                    hover:text-white 
                    cursor-pointer ${
                        isActive.isActive ? "text-white" : ""
                    }`,
                    className,
                ]);
            }}
        >
            {children}
            <span className="text-4xl">
        <Icon/>
      </span>
            <span>{name}</span>
        </NavLink>
    );
};

export default SideBarItem;
