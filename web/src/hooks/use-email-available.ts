import { useSubmit } from "./use-submit";
import { useEffect } from "react";

export const useEmailAvailable = (email: string): boolean => {
    const { data, submit } = useSubmit<{ isAvailable: boolean }, { value: string }>({
        url: "/v1/ark/users/email-available"
    });

    useEffect(() => {
        if(email !== "" && email.includes("@") && email.includes(".")) {
            submit({ value: email });
        }
    }, [email]);

    return data?.isAvailable !== false;
};
