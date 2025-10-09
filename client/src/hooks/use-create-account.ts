import React, { useEffect } from "react";
import {useSubmit} from "../hooks";

const useCreateAccount = () => {
    const { state, error, submit } = useSubmit({ url: `/oauth2/register`});


    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        const email = formData.get("email")?.toString() ?? "";
        const username = formData.get("username")?.toString() ?? "";
        const password = formData.get("password")?.toString() ?? "";
        const firstName = formData.get("firstName")?.toString() ?? "";
        const lastName = formData.get("lastName")?.toString() ?? "";

        submit({username, email, password, firstName, lastName, isAdmin: false})
    }

    useEffect(() => {
        if (state === "success") {
            window.location.href = "/login";
        }
    }, [state])

    return { state, error, onSubmit }

}

export default useCreateAccount