import { FC } from "react";
import { Layout } from "../../components";
import Card from "../../components/core/card/card.tsx";
import { useSearchParams } from "react-router";

const ErrorPage: FC = () => {
    const [searchParams] = useSearchParams();
    const errorMessage = searchParams.get("error");

    return (
        <Layout>
            <Card
                title="Error"
                description="An error has occured during authorization"
            >
                <div className="text-muted-foreground border border-zinc-600 bg-zinc-700 rounded-lg p-4 text-center w-full">
                    {errorMessage}
                </div>
            </Card>
        </Layout>
    );
};

export default ErrorPage;
