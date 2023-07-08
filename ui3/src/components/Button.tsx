import React, {FC} from "react";
import {IconType} from "react-icons";
import {twJoin} from "tailwind-merge";

interface ButtonProps {
    className?: string;
    text?: string;
    icon?: IconType;
    onClickFun: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
}

export const Button: FC<ButtonProps> = ({
                                            className,
                                            icon: Icon,
                                            text,
                                            onClickFun,
                                        }) => {
    return (
        <button
            type="button"
            onClick={onClickFun}
            className={twJoin([
                "bg-amber-400 p-2 flex justify-center items-center",
                className,
            ])}
        >
            {Icon === undefined ? "" : <Icon className="text-3xl"/>}
            {text}
        </button>
    );
};
