
interface HeaderProps {
    text: string;
}

export const SubHeader = ({text}:HeaderProps) => {
    return(
        <p className=
        {
            `text-pf-gray-100 font-bold text-2xl mt-3`
        }>
            {text}
        </p>
    );  
}

export const Content = ({text}:HeaderProps) => {
    return(
        <p className=
        {
            `text-pf-accent-2 font-bold text-2xl mt-3`
        }>
            {text}
        </p>
    );  
}

