import { FC, useEffect } from "react";

const IndexPage: FC = () => {
    useEffect(() => {
        window.location.href = "/login";
    }, []);
    return null;
};

export default IndexPage;
