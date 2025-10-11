import { useEffect } from "react";

export const useTheme = () => {
    useEffect(() => {
        document.documentElement.classList.toggle(
            "dark",
            window.matchMedia("(prefers-color-scheme: dark)").matches
        );
    }, []);
};
