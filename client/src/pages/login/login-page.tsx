import {FC} from 'react';
import {Button, ErrorBanner, Layout} from "../../components";
import {useLogin} from "../../hooks";
import Card from "../../components/core/card/card.tsx";

const LoginPage: FC = () => {
    const { state, error, onSubmit } = useLogin();

    return <Layout>
        <Card title='Sign in' description='Use your Eden account'>
            {
                error && <ErrorBanner title="An error has occured">{error.message}</ErrorBanner>
            }
            <form className='flex flex-col gap-8 min-w-72' onSubmit={onSubmit}>
                <div className="flex flex-col gap-4">
                    <div className='flex flex-col gap-1.5'>
                        <label htmlFor='email'>E-mail</label>
                        <input id='email' name="email" type="email" placeholder="Enter e-mail" required/>
                    </div>
                    <div className='flex flex-col gap-1.5'>
                        <label htmlFor='password'>Password</label>
                        <input id='password' name="password" type="password" placeholder="Enter password" required/>
                        <a href="/forgot-password" className="link mt-0.5">Forgot password ?</a>
                    </div>
                </div>

                <div className="flex items-center justify-between gap-8">
                    <a href="/create-account">
                        <Button type='button' variant="secondary">Create account</Button>
                    </a>
                    <Button variant="primary" disabled={state === 'pending'}>Log in</Button>
                </div>
            </form>
        </Card>
    </Layout>
};

export default LoginPage;