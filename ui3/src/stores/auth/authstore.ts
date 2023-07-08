import {create} from 'zustand';
import {client} from '../../utils/client';
import {User, UserCreateResponseDto} from '../../models/user';
import {AUTH_STORE_KEY} from '../../constants/auth';

interface State {
    isLoggedIn: boolean;
    userid: string;
    authToken: string;
    user: User | undefined;
}

interface Action {
    login: (email: string, password: string) => Promise<void>;
    Register: (email: string, username: string, password: string, role: number) => Promise<void>;
    logout: () => void;
    checkLocalAuth: () => Promise<void>;
}

export const useAuthStore = create<State & Action>()((set, get) => ({
    isLoggedIn: false,
    userid: "",
    authToken: "",
    user: undefined,
    login: async (email: string, password: string) => {
        const data = await SignIn(email, password);
        set(() => ({authToken: data.authToken, userid: data.userid, isLoggedIn: data.isLoggedIn}));
    },
    Register: async (email: string, username: string, password: string, role: number) => {
        const data = await SignUp(email, username, password, role);
        console.log(data);

        set(() => {
            return {
                authToken: data.authToken,
                isLoggedIn: true,
                userid: data.userid,
                user: data.user
            };
        });
    },
    logout: () => {
        set(() => ({isLoggedIn: false, authToken: "", userid: "", user: undefined}));
        localStorage.removeItem(AUTH_STORE_KEY);

    },
    checkLocalAuth: async () => {
        const data = await checkLocalAuth(get());
        console.log(data);

        set(() => ({authToken: data?.authToken, isLoggedIn: data?.isLoggedIn, userid: data?.userid, user: data?.user}));
    }
}));

const checkLocalAuth = async (state: (State & Action)) => {
    if (!state.isLoggedIn && state.authToken === "" && state.userid === "") {
        const storedData = localStorage.getItem(AUTH_STORE_KEY);
        const data = JSON.parse(storedData === null ? "" : storedData);
        console.log(data);

        if (data !== null) {
            if (await checkTokenValidity(data.token)) {
                console.log(data);
                state.authToken = data.token;
                state.isLoggedIn = true;
                state.userid = data.user.id;

                const returnData: State = {
                    authToken: data.token,
                    isLoggedIn: true,
                    userid: data.user.id,
                    user: data.user
                };

                return returnData;
            }
            localStorage.removeItem(AUTH_STORE_KEY);
        }
    }
};

const checkTokenValidity = async (token: string): Promise<boolean> => {
    console.log(token);

    const response = await client.post("/user/isValid", {}, {
        headers: {
            Authorization: token
        }
    });
    console.log(response);

    if (response) {
        console.log("auth sucess");

        return true;
    }
    return false;
};

const SignIn = async (email: string, password: string): Promise<State> => {
    const response = await client.post("/user/signin", {
        email,
        password,
    });
    if (response.status === 202) {
        const data: UserCreateResponseDto = response.data;

        const returnData: State = {
            authToken: data.token,
            isLoggedIn: true,
            userid: data.user.id,
            user: data.user
        };
        localStorage.setItem(AUTH_STORE_KEY, JSON.stringify(data));
        return returnData;
    }
    return {
        authToken: "",
        isLoggedIn: false,
        userid: "",
        user: undefined
    };
};

const SignUp = async (email: string, username: string, password: string, role = 1): Promise<State> => {
    localStorage.removeItem(AUTH_STORE_KEY);
    const response = await client.post("/user/signup", {
        name: username,
        email,
        password,
        role
    });
    if (response.status === 202) {
        const data: UserCreateResponseDto = response.data;

        const returnData: State = {
            authToken: data.token,
            isLoggedIn: true,
            userid: data.user.id,
            user: data.user
        };

        localStorage.setItem(AUTH_STORE_KEY, JSON.stringify(data));
        return returnData;
    }
    return {
        authToken: "",
        isLoggedIn: false,
        userid: "",
        user: undefined
    };
};