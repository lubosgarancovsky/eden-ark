import { FC } from "react";
import { Button, ErrorBanner, Input, Layout, Progress } from "../../components";
import Card from "../../components/core/card/card.tsx";
import { useChangePassword } from "../../hooks";

const ResetPasswordPage: FC = () => {
    const {
        password,
        password2,
        onPasswordChange,
        timeLeft,
        percentage,
        state,
        error,
        onSubmit
    } = useChangePassword();

    const mismatch = password !== password2 && password2 != "";
    const accent = mismatch ? "error" : password2 != "" ? "success" : undefined;

    return (
        <Layout>
            <Card title="Reset password" description="Create new password">
                {state === "success" && (
                    <div className="flex flex-col gap-8">
                        <ErrorBanner title="Success" variant="success">
                            Your password was successfully changed.
                        </ErrorBanner>

                        <div className="flex items-center justify-end gap-8">
                            <a href="/login">
                                <Button variant="primary">Sign in</Button>
                            </a>
                        </div>
                    </div>
                )}
                {state !== "success" && (
                    <form
                        className="flex flex-col gap-8 min-w-72"
                        onSubmit={onSubmit}
                    >
                        <div className="flex flex-col gap-4">
                            <div className="flex flex-col gap-1.5">
                                <Input
                                    label="Password"
                                    name="password"
                                    type="password"
                                    value={password}
                                    onChange={onPasswordChange("password")}
                                    required
                                />
                            </div>
                            <div className="flex flex-col gap-1.5">
                                <Input
                                    label="Repeat password"
                                    name="password2"
                                    type="password"
                                    value={password2}
                                    accent={accent}
                                    onChange={onPasswordChange("password2")}
                                    required
                                />
                            </div>
                        </div>

                        <div>
                            <div className="flex gap-4 items-center">
                                {percentage  > 0 ? <Progress value={percentage} /> : (
                                    <p className="text-red-500 text-sm w-full">
                                        Session timed out
                                    </p>
                                )}
                                <div>{timeLeft}</div>
                            </div>
                        </div>


                        <div className="flex items-center justify-between gap-8">
                            <a href="/login" className="-ml-4">
                                <Button type="button" variant="secondary">
                                    Back to login
                                </Button>
                            </a>
                            <div className="flex items-center gap-4">
                                {error && (
                                    <p className="text-red-500 text-sm">
                                        {error?.message}
                                    </p>
                                )}

                                <Button
                                    variant="primary"
                                    disabled={state === "pending" || mismatch || percentage === 0}
                                >
                                    Change password
                                </Button>
                            </div>
                        </div>
                    </form>
                )}
            </Card>
        </Layout>
    );
};

export default ResetPasswordPage;
