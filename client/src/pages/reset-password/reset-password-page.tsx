import {FC} from 'react';
import {Button, ErrorBanner, Input, Layout} from "../../components";
import Card from "../../components/core/card/card.tsx";
import {useChangePassword} from "../../hooks";
import {cn} from "../../lib";


const ResetPasswordPage: FC = () => {
    const {password, password2, onPasswordChange, state, error, onSubmit} = useChangePassword();


    return (
        <Layout>
            <Card title="Reset password" description="Create new password">
                {
                    error && <ErrorBanner title="An error has occured">{error.message}</ErrorBanner>
                }
                {state === 'success' &&
                    <div className="flex flex-col gap-8">
                        <ErrorBanner title="Success" variant='success'>
                            Your password was successfully changed.
                        </ErrorBanner>

                        <div className="flex items-center justify-end gap-8">
                            <a href="/login">
                            <Button variant="primary">Sign in</Button>
                            </a>
                        </div>
                    </div>
                }
                {state !== 'success' && <form className='flex flex-col gap-8 min-w-72' onSubmit={onSubmit}>
                    <div className="flex flex-col gap-4">
                        <div className='flex flex-col gap-1.5'>
                            <Input label="Password" name="password" type="password" value={password}
                                   onChange={onPasswordChange('password')} required/>
                        </div>
                        <div className='flex flex-col gap-1.5'>
                            <Input
                                label="Repeat password"
                                name="password2"
                                type="password"
                                value={password2}
                                className={cn(
                                    {
                                        "!border-green-500 !focus-visible:ring-green-500 !text-green-500": password === password2 && password2 != "",
                                        "!border-red-500 !focus-visible:ring-red-500 !text-red-500": password !== password2 && password2 != "",
                                    }
                                )}
                                onChange={onPasswordChange('password2')}
                                required/>
                        </div>
                    </div>

                    <div className="flex items-center justify-between gap-8">
                        <a href="/login">
                            <Button type='button' variant="secondary">Back to login</Button>
                        </a>
                        <Button variant="primary"
                                disabled={state === 'pending' || (password != password2 && password2 != "")}>Change
                            password</Button>
                    </div>
                </form>}
            </Card>
        </Layout>
    );
};

export default ResetPasswordPage;