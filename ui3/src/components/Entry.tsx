import React, {FC} from "react";
import {twMerge} from "tailwind-merge";

interface EntryProps {
    className?: string;
    placeHolder?: string;
    value?: string;
    onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

export const Entry: FC<EntryProps> = ({
                                          className,
                                          placeHolder,
                                          value,
                                          onChange,
                                      }) => {
    return (
        <input
            value={value}
            className={twMerge([
                "form-input placeholder:text-gray-800 rounded-sm bg-gray-500",
                className,
            ])}
            placeholder={placeHolder}
            onChange={onChange}
        />
    );
};
