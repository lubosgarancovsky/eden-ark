import { FC } from 'react';
import { Button, Input, Layout } from "../../components";
import Card from "../../components/core/card/card.tsx";

const CreateAccountPage: FC = () => {
    const onSubmit = () => {}

    return (
        <Layout>
            <Card title="Create account" description="Create new Eden account">
                <form className='flex flex-col gap-8 min-w-72' onSubmit={onSubmit}>
                    <div className="flex flex-col gap-4">
                            <Input label="First name" name="firstName" type="text" required/>
                            <Input label="Last name" name="lastName" type="text" required/>
                            <Input label="E-mail" name="email" type="email" required/>
                            <Input label="Username" name="username" type="text" required/>
                            <Input label="Password" name="password" type="password" required/>
                            <Input label="Repeat password" name="password2" type="password" required/>
                    </div>

                    <div className="flex items-center justify-between gap-8">
                        <a href="/login" className="-ml-4">
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