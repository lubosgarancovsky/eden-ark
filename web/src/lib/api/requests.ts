import axios from "axios";
import type {SignUpRequest} from "@/lib/api";

export const createLoginSession = async (username: string, password: string, returnTo: string) => {
    await axios.post(`/oauth2/login?returnTo=${returnTo}`, { username, password });
}

export const isUsernameAvailable = async (username: string): Promise<boolean> => {
    const result = await axios.get(`/v1/ark/users/username-available/${username}`);
    return result.data.isAvailable;
}

export const isEmailAvailable = async (email: string): Promise<boolean> => {
    const result = await axios.get(`/v1/ark/users/email-available/${email}`);
    return result.data.isAvailable;
}

export const signup = async (data: SignUpRequest) => {
    const result = await axios.post(`/v1/ark/users`, data);
    return result.data;
}