
interface ListitemProps {
    accountname: string;
    balance: number;
    description: string;
}


export const Listitem = ({accountname,balance,description}: ListitemProps) =>{
    return(
        <div className='flex justify-between items-center'>
            <div className='flex justify-between items-center'>
                <div className='text-pf-gray-100 flex justify-center items-center rounded-full bg-[#F77F00] w-14 h-14'>
                    <p className='text-[24px]'>{accountname[0]}</p>
                </div>
                <div className='ml-5'>
                    <p className='text-pf-gray-900 font-bold text-2xl'>{accountname}</p>
                    <p className='text-pf-gray-900 font-normal text-2xl'>{description}</p>
                </div>
            </div>
            <div className='flex justify-between items-center'>
                <p className='text-pf-gray-900 font-bold text-2xl'>Balance: </p>
                <p className='text-pf-gray-900 font-bold text-2xl'>{balance.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
            </div>
        </div>
    )
}