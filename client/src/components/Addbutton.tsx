
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlus } from '@fortawesome/free-solid-svg-icons';

interface AddButtonProps {
    text: string;
    func?: () => void | ((arg: any) => void);
}


export const AddButton = ({text,func}:AddButtonProps) => {

    return (
        <div className='flex justify-between items-center'>
            <p className='text-[24px] mr-3'>{text}</p>
            <div 
                className='text-pf-gray-100 flex justify-center items-center rounded-full bg-[#F77F00] w-14 h-14 hover:bg-pf-accent-2 cursor-pointer'
                onClick={func}
            >
                <FontAwesomeIcon 
                    icon={faPlus} 
                    size='1x'
                />
            </div>
        </div>
    )
}