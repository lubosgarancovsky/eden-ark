import { FC, useEffect } from "react";

const LogoutPage: FC = () => {
    useEffect(() => {
        window.location.href = "/oauth2/logout";
    }, []);
    return null;
};

export default LogoutPage;
