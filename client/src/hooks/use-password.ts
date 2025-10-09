import { ChangeEvent, useState } from "react";

export const usePassword = () => {
    const [password, setPassword] = useState('');
    const [password2, setPassword2] = useState('');

    const handlePasswordChange = (variant: 'password' | 'repeat-password') => (e: ChangeEvent<HTMLInputElement>) => {
        if (variant === 'password') {
            setPassword(e.target.value);
            return
        }

        setPassword2(e.target.value);
    }

    const mismatch = password !== password2 && password2 != "";
    const accent: "success" | "error" | undefined = mismatch ? "error" : password2 != "" ? "success" : undefined;
    return { password, password2, mismatch, accent, handlePasswordChange };
};