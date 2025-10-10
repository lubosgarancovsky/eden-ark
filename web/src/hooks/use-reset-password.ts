import React, {useState} from "react";
import {useSubmit} from "./use-submit";

export const useResetPassword = () => {
    const [email, setEmail] = useState('');

    const { state, error, submit } = useSubmit({ url: "/v1/ark/users/request-reset-password" })

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        const email = formData.get("email")?.toString() ?? "";
        setEmail(email);
        submit({ email })
    }

    const resend = () => {
        submit({ email })
    }

    return { onSubmit, email, state, error, resend }
}