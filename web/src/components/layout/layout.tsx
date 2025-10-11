import { FC, ReactNode } from "react";
import { useTheme } from "../../hooks/use-theme.ts";

type Props = {
    children: ReactNode;
};

const Layout: FC<Props> = ({ children }) => {
    useTheme();

    return (
        <div className="relative flex flex-col p-32 gap-8 items-center min-h-screen">
            {children}
        </div>
    );
};

export default Layout;
