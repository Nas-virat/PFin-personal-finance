import { faArrowRight } from "@fortawesome/free-solid-svg-icons/faArrowRight"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"

interface HeaderCardProps {
    text: string;
    func?: () => void | ((arg: any) => void);
}

export const HeaderCard = ({text,func}:HeaderCardProps) => {

    return (
        <div className='flex justify-between'
            onClick={func}
        >
            <p className="text-pf-gray-100 font-bold text-3xl">{text}</p>
            <FontAwesomeIcon
                    icon={faArrowRight}
                    size='3x'
                    className='text-pf-gray-100 hover:text-pf-accent-2'
            />
        </div>
    )
}