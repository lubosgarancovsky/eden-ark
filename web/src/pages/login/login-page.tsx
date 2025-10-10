import { FC } from "react";
import { Button, Input, Layout } from "../../components";
import { useLogin } from "../../hooks";
import Card from "../../components/core/card/card.tsx";
import { Check } from "lucide-react";

const LoginPage: FC = () => {
    const { state, error, returnTo } = useLogin();

    if (state === "success") {
        return (
            <Layout>
                <Card title="Success" description="You are now signed in">
                    <div className="border border-green-500 bg-green-500/10 rounded-full mx-auto w-32 h-32 text-green-500 flex items-center justify-center">
                        <Check size={64} />
                    </div>
                </Card>
            </Layout>
        );
    }

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
                                label="E-mail"
                                name="email"
                                type="email"
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
                        <p className="text-red-500 text-sm">{error?.message}</p>
                    )}

                    <div className="flex items-center justify-between gap-8">
                        <a href="/create-account" className="-ml-4">
                            <Button type="button" variant="secondary">
                                Create account
                            </Button>
                        </a>

                        <Button variant="primary">Log in</Button>
                    </div>
                </form>
            </Card>
        </Layout>
    );
};

export default LoginPage;
