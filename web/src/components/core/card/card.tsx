import { FC, ReactNode } from "react";
import { ChevronLeft } from "lucide-react";

type Props = {
    title: string;
    description: string;
    backLink?: { label: string; href: string };
    children?: ReactNode;
};

const Card: FC<Props> = ({ title, description, backLink, children }) => {
    return (
        <div className="flex flex-col gap-8 rounded-lg p-8 dark:py-0 bg-card shadow relative lg:w-[28rem]">
            {backLink && (
                <a
                    href={backLink.href}
                    className="flex gap-1 items-center w-fit hover:underline text-sm"
                >
                    <ChevronLeft size={16} />
                    {backLink.label}
                </a>
            )}
            <div className="flex flex-col gap-0.5 ">
                <h1>{title}</h1>
                <p className="text-muted-foreground text-sm">{description}</p>
            </div>
            <div>{children}</div>
        </div>
    );
};

export default Card;
