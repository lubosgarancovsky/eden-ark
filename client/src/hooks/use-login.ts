import React from "react";
import {useSubmit} from "../hooks";
import {useSearchParams} from "react-router";

const useLogin = () => {
    const [urlSearchParams] = useSearchParams();
    const returnTo = urlSearchParams.get("returnTo") as string
    const { state, error, submit } = useSubmit({ url: `/oauth2/login?returnTo=${returnTo}`});


    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        const email = formData.get("email")?.toString() ?? "";
        const password = formData.get("password")?.toString() ?? "";

        submit({email, password})
    }

    return { state, error, onSubmit }
}

export default useLogin