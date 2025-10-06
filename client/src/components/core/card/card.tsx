import {FC, ReactNode} from 'react';

type Props = {
    title: string;
    description: string;
    children?: ReactNode;
};

const Card: FC<Props> = ({ title, description, children}) => {
 return (
     <div className="bg-neutral-950 p-8 rounded-4xl flex flex-col gap-8 md:w-xl w-full border border-neutral-900 shadow-black shadow-2xl">
         <div className="w-16 h-16 bg-blue-700 rounded-full"></div>
         <div className="flex flex-col gap-2">
             <h1>{title}</h1>
             <p className="text-neutral-400">{description}</p>
         </div>
         {children}
     </div>
 );
};

export default Card;