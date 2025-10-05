import React from "react";

const sendLoginCredentials = async (email: string, password: string, returnTo: string) => {
    try {
        const result = await fetch(`/oauth2/login?returnTo=${returnTo}`, {
            method: "POST",
            body: JSON.stringify({email, password}),
        })

        return result.ok;
    } catch(e) {
        console.error(e)
        return false;
    }
}

const useLogin = (returnTo: string) => {
    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        const email = formData.get("email")?.toString() ?? "";
        const password = formData.get("password")?.toString() ?? "";

        await sendLoginCredentials(email, password, returnTo);
    }

    return { onSubmit }
}

export default useLogin