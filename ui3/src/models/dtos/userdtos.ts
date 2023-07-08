import type {User} from "../user";

export interface SigUpDto {
    name: string;
    email: string;
    password: string;
    role: number;
}

export interface SigInDto {
    email: string;
    password: string;
}

export interface SignUpRespDto {
    user: User;
    token: string;
}