import type {BaseModel} from "./base";

export interface UserCreateResponseDto {
    user: User;
    token: string;
}

export interface User extends BaseModel {
    name: string;
    email: string;
    role: number;
}

export interface UserCreateDto {
    name: string;
    password: string;
    email: string;
}