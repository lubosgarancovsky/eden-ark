import {FC} from 'react';
import { Button, Input, Layout } from "../../components";
import {useLogin} from "../../hooks";
import Card from "../../components/core/card/card.tsx";

const LoginPage: FC = () => {
    const { state, error, onSubmit } = useLogin();

    return <Layout>
        <Card title='Sign in' description='Use your Eden account'>
            <form className='flex flex-col gap-8 min-w-72' onSubmit={onSubmit}>
                <div className="flex flex-col gap-4">
                    <div className='flex flex-col gap-1.5'>
                        <Input label="E-mail" name="email" type="email" required/>
                    </div>
                    <div className='flex flex-col gap-1.5'>
                        <Input label="Password" name="password" type="password" required/>
                        <a href="/forgot-password" className="link mt-0.5">Forgot password ?</a>
                    </div>
                </div>

                <div className="flex items-center justify-between gap-8">
                    <a href="/create-account" className="-ml-4">
                        <Button type='button' variant="secondary">Create account</Button>
                    </a>
                    <div className="flex items-center gap-4">
                        {error && <p className="text-red-500 text-sm">{error?.message}</p>}
                        <Button variant="primary" disabled={state === 'pending'}>Log in</Button>
                    </div>
                </div>
            </form>
        </Card>
    </Layout>
};

export default LoginPage;