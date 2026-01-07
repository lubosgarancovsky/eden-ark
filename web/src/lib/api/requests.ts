import { httpClient } from "@/lib/axios";

export const isUsernameAvailable = async (username: string): Promise<boolean> => {
    const result = await httpClient.get(`/v1/ark/users/username-available/${username}`);
    return result.data.isAvailable;
}

export const requestPasswordReset = async (email: string) => {
    await httpClient.post(`/v1/ark/users/request-reset-password`, {email})
}

export const changePassword = async (password: string, token: string) => {
    await httpClient.post(`/v1/ark/users/reset-password`, {password, token})
}

