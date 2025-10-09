import { ChangeEvent, FC, useState } from "react";
import { Button, Input, Layout } from "../../components";
import Card from "../../components/core/card/card.tsx";
import { useEmailAvailable } from "../../hooks/use-email-available.ts";
import useDebounce from "../../hooks/use-debounce.ts";
import { useUsernameAvailable } from "../../hooks/use-username-available.ts";
import useCreateAccount from "../../hooks/use-create-account.ts";
import { usePassword } from "../../hooks/use-password.ts";

const CreateAccountPage: FC = () => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const { password, password2, handlePasswordChange, accent, mismatch } =
        usePassword();

    const emailValue = useDebounce(email);
    const usernameValue = useDebounce(username);

    const isEmailAvailable = useEmailAvailable(emailValue);
    const isUsernameAvailable = useUsernameAvailable(usernameValue);

    const { onSubmit, error } = useCreateAccount();

    const handleEmail = (e: ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value);
    };

    const handleUsername = (e: ChangeEvent<HTMLInputElement>) => {
        setUsername(e.target.value);
    };

    return (
        <Layout>
            <Card title="Create account" description="Create new Eden account">
                <form
                    className="flex flex-col gap-8 min-w-72"
                    onSubmit={onSubmit}
                >
                    <div className="flex flex-col gap-4">
                        <div className="flex flex-col lg:flex-row gap-4">
                            <Input
                                label="First name"
                                name="firstName"
                                type="text"
                                required
                            />
                            <Input
                                label="Last name"
                                name="lastName"
                                type="text"
                                required
                            />
                        </div>
                        <Input
                            label="E-mail"
                            name="email"
                            type="email"
                            value={email}
                            onChange={(e) => handleEmail(e)}
                            error={
                                !isEmailAvailable
                                    ? "This e-mail is already used"
                                    : undefined
                            }
                            required
                        />
                        <Input
                            label="Username"
                            name="username"
                            type="text"
                            value={username}
                            onChange={(e) => handleUsername(e)}
                            error={
                                !isUsernameAvailable
                                    ? "This username is already taken"
                                    : undefined
                            }
                            required
                        />
                        <Input
                            label="Password"
                            name="password"
                            type="password"
                            value={password}
                            onChange={handlePasswordChange("password")}
                            required
                        />
                        <Input
                            label="Repeat password"
                            name="password2"
                            type="password"
                            value={password2}
                            onChange={handlePasswordChange("repeat-password")}
                            accent={accent}
                            required
                        />
                    </div>

                    {error && (
                        <p className="text-red-500 text-sm">{error.message}</p>
                    )}

                    <div className="flex items-center justify-between gap-8">
                        <a href="/login" className="-ml-4">
                            <Button type="button" variant="secondary">
                                Back to login
                            </Button>
                        </a>
                        <Button variant="primary" disabled={mismatch}>
                            Create account
                        </Button>
                    </div>
                </form>
            </Card>
        </Layout>
    );
};

export default CreateAccountPage;
