import React, {ChangeEvent, useState} from "react";
import {useSearchParams} from "react-router";
import {useSubmit} from "./use-submit";
import { useCountdown } from "./use-countdown.ts";

const URI = '/v1/ark/users/reset-password';

export const useChangePassword = () => {
    const [params] = useSearchParams();

    const token = params.get("token");
    const expiresAt = params.get("expiresAt");

    const { timeLeft, percentage} = useCountdown(expiresAt ?? "");

    const { state, error, submit } = useSubmit({
        url: URI
    })

    const [password, setPassword] = useState('');
    const [password2, setPassword2] = useState('');

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        const password = formData.get('password');
        const password2 = formData.get('password2');

        if (password === password2) {
            submit({ token, password });
        }
    }

    const onPasswordChange = (name: 'password' | 'password2') => (e: ChangeEvent<HTMLInputElement>) => {
        if (name === 'password') {
            setPassword(e.target.value);
        } else {
            setPassword2(e.target.value);
        }
    }

    return { state, error, password, password2, timeLeft, percentage, expiresAt, onSubmit, onPasswordChange }
}