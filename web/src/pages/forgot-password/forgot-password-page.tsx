import { FC } from "react";
import { Button, Input, Layout } from "../../components";
import Card from "../../components/core/card/card.tsx";
import { useResetPassword } from "../../hooks";
import { Mail } from "lucide-react";

const ForgotPasswordPage: FC = () => {
    const { error, state, email, onSubmit, resend } = useResetPassword();

    return (
        <Layout>
            <Card
                title="Reset password"
                description="Enter your e-mail to reset password"
                backLink={{ label: "Back to login", href: "/login" }}
            >
                {state === "success" && (
                    <div className="flex flex-col gap-4">
                        <p>
                            An e-mail with a verification link was sent to your
                            e-mail address:
                        </p>
                        <div className="p-4 bg-black/20 rounded-lg text-muted-foreground text-center">
                            {email}
                        </div>

                        <div className="flex items-center justify-between gap-8">
                            <a href="/login" className="-ml-4">
                                <Button type="button" variant="secondary">
                                    Back to login
                                </Button>
                            </a>
                            <Button variant="primary" onClick={resend}>
                                Resend e-mail
                            </Button>
                        </div>
                    </div>
                )}
                {state !== "success" && (
                    <form
                        className="flex flex-col gap-8 min-w-72"
                        onSubmit={onSubmit}
                    >
                        <div className="flex flex-col gap-4">
                            <Input
                                label="E-mail"
                                id="email"
                                name="email"
                                type="email"
                                error={error?.message}
                                required
                            />
                        </div>

                        <p className="text-muted-foreground text-sm">
                            A verification link will be sent to your e-mail
                        </p>

                        <div className="flex flex-col gap-2">
                            <Button
                                variant="primary"
                                disabled={state === "pending"}
                            >
                                <Mail size={16} />
                                Send verification e-mail
                            </Button>
                        </div>
                    </form>
                )}
            </Card>
        </Layout>
    );
};

export default ForgotPasswordPage;
