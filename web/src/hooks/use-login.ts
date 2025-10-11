import { useSearchParams } from "react-router";

const useLogin = () => {
    const [urlSearchParams] = useSearchParams();
    const returnTo = urlSearchParams.get("returnTo") as string;
    const errorMsg = urlSearchParams.get("error") as string;
    return { error: errorMsg, returnTo };
};

export default useLogin;
