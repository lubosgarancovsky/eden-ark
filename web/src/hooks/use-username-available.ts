import { useSubmit } from "./use-submit";
import { useEffect } from "react";

export const useUsernameAvailable = (username: string): boolean => {
    const { data, submit } = useSubmit<{ isAvailable: boolean }, { value: string }>({
        url: "/v1/ark/users/username-available"
    });

    useEffect(() => {
        if(username !== "" ) {
            submit({ value: username });
        }
    }, [username]);

    return data?.isAvailable !== false;
};
