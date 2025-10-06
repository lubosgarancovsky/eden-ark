import { FC } from 'react';
import {Button, Layout} from "../../components";
import Card from "../../components/core/card/card.tsx";

const CreateAccountPage: FC = () => {
    const onSubmit = () => {}

    return (
        <Layout>
            <Card title="Create account" description="Create new Eden account">
                <form className='flex flex-col gap-8 min-w-72' onSubmit={onSubmit}>
                    <div className="flex flex-col gap-4">
                        <div className='flex flex-col gap-1.5'>
                            <label htmlFor='firstName'>First name</label>
                            <input id='firstName' name="firstName" type="text" required/>
                        </div>
                        <div className='flex flex-col gap-1.5'>
                            <label htmlFor='lastName'>Last name</label>
                            <input id='lastName' name="lastName" type="text" required/>
                        </div>
                        <div className='flex flex-col gap-1.5'>
                            <label htmlFor='email'>E-mail</label>
                            <input id='email' name="email" type="email" required/>
                        </div>
                        <div className='flex flex-col gap-1.5'>
                            <label htmlFor='username'>Username</label>
                            <input id='username' name="username" type="text" required/>
                        </div>
                        <div className='flex flex-col gap-1.5'>
                            <label htmlFor='password'>Password</label>
                            <input id='password' name="password" type="password" required/>
                        </div>
                        <div className='flex flex-col gap-1.5'>
                            <label htmlFor='password2'>Repeat password</label>
                            <input id='password2' name="password2" type="password" required/>
                        </div>
                    </div>

                    <div className="flex items-center justify-between gap-8">
                        <a href="/login">
                            <Button type='button' variant="secondary">Back to login</Button>
                        </a>
                        <Button variant="primary">Create account</Button>
                    </div>
                </form>
            </Card>
        </Layout>
    );
};

export default CreateAccountPage;