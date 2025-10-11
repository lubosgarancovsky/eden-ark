import { FC } from "react";
import { Button, Input, Layout } from "../../components";
import { useLogin } from "../../hooks";
import Card from "../../components/core/card/card.tsx";
import { ArrowRight } from "lucide-react";

const LoginPage: FC = () => {
    const { error, returnTo } = useLogin();

    return (
        <Layout>
            <Card title="Sign in" description="Use your Eden account">
                <form
                    id="loginForm"
                    action="/oauth2/login"
                    method="POST"
                    className="flex flex-col gap-8 min-w-72"
                >
                    <div className="flex flex-col gap-4">
                        <input type="hidden" name="returnTo" value={returnTo} />
                        <div className="flex flex-col gap-1.5">
                            <Input
                                label="E-mail or username"
                                name="username"
                                type="username"
                                required
                            />
                        </div>
                        <div className="flex flex-col gap-1.5">
                            <Input
                                label="Password"
                                name="password"
                                type="password"
                                required
                            />
                            <a href="/forgot-password" className="link mt-0.5">
                                Forgot password ?
                            </a>
                        </div>
                    </div>

                    {error && (
                        <p className="text-red-500 text-sm">{error}</p>
                    )}

                    <Button variant="primary">
                        Log in <ArrowRight size={16} />
                    </Button>
                </form>
            </Card>
            <div className="text-sm">
                Don't have an account yet?{" "}
                <a href="/create-account" className="link">
                    Sign up
                </a>
            </div>
        </Layout>
    );
};

export default LoginPage;
