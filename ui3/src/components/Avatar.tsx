import {FC, useMemo} from "react";
import {twJoin} from "tailwind-merge";

interface AvatarProp {
    AvatarName: string;
    fontStyle?: string;
    className?: string;
}

export const Avatar: FC<AvatarProp> = ({
                                           AvatarName,
                                           fontStyle,
                                           className,
                                       }) => {
    function getAvatarCharacter(name: string) {
        if (name === "" || name === undefined) {
            return;
        }
        return name.charAt(0);
    }

    //Todo colors set and choose
    function randomCol() {
        const randVal = Math.random() * (255);
        return Math.round(randVal);
    }

    function generateRandomColor() {
        const r = randomCol();
        const g = randomCol();
        const b = randomCol();
        return `rgba(${r},${g},${b},0.8)`;
    }


    const avatarColor = useMemo(() => generateRandomColor(),[Avatar]) //Todo Find fix


    return (
        <div
            className={twJoin([
                `rounded-full border-2 border-white w-12 h-12 flex justify-center  items-center p-1`,
                className,
            ])}
            style={{backgroundColor: `${avatarColor}`}}
        >
            <h2 className={twJoin(["uppercase font-semibold text-3xl", fontStyle])}>
                {getAvatarCharacter(AvatarName)}
            </h2>
        </div>
    );
};
