import { FC, DetailedHTMLProps, ButtonHTMLAttributes } from "react";
import { cn } from "../../../lib";

type Props = DetailedHTMLProps<
    ButtonHTMLAttributes<HTMLButtonElement>,
    HTMLButtonElement
> & {
    variant?: "primary" | "secondary";
};

const Button: FC<Props> = ({
    children,
    variant = "primary",
    className,
    ...props
}) => {
    return (
        <button
            className={cn(
                "text-sm px-4 py-2 rounded cursor-pointer flex items-center justify-center gap-2",
                {
                    "bg-primary hover:bg-primary/80 disabled:bg-primary/10 disabled:text-foreground/50 text-white":
                        variant === "primary",
                    "bg-transparent hover:bg-black/20 text-neutral-100":
                        variant === "secondary"
                },
                className
            )}
            {...props}
        >
            {children}
        </button>
    );
};

export default Button;
