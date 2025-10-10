import React, { useEffect } from "react";
import {useSearchParams} from "react-router";
import {useSubmit} from "./use-submit";
import { useCountdown } from "./use-countdown.ts";
import { usePassword } from "./use-password.ts";

const URI = '/v1/ark/users/reset-password';

export const useChangePassword = () => {
    const [params] = useSearchParams();
    const {password, password2, handlePasswordChange, accent, mismatch} = usePassword();

    const token = params.get("token");
    const expiresAt = params.get("expiresAt");

    const { timeLeft, percentage} = useCountdown(expiresAt ?? "");

    const { state, error, submit } = useSubmit({
        url: URI
    })

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        const password = formData.get('password');
        const password2 = formData.get('password2');

        if (password === password2) {
            submit({ token, password });
        }
    }

    useEffect(() => {
        if (state === "success") {
            window.location.href = "/login";
        }
    }, [state])

    return { state, error, password, password2, accent, mismatch, timeLeft, percentage, expiresAt, onSubmit, handlePasswordChange}
}