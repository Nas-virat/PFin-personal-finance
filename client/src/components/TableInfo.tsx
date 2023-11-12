import * as React from 'react';


export const TableInfo = ({columns,data,total}:TableInfoProps)  => {

  const emptyCells = [];
  if(total !== undefined){
    for (let i = 0; i < columns.length - 2; i++) {
      emptyCells.push(
        <td key={i} className="px-6 py-4 text-xl font-semibold"></td>
      );
    }
  }

  return (
    <>
      <table className="mx-4 mt-7 border-collapse min-w-[90%] rounded-xl overflow-hidden">
        <thead>
          <tr className=" bg-pf-secondary-2 text-pf-gray-100">
            {
              columns.map((key,column) => (
                <th key={key} className="px-6 py-4 text-2xl font-semibold">{column}</th>
              ))
            }
          </tr>
        </thead>
        <tbody>  
          {
            data.map((item,index) => (
              <tr className="bg-pf-gray-100 text-pf-gray-900 text-center" key={index}>
                {
                  Object.keys(item).map((key,index) => (
                    <td key={'object-'+index} className="px-6 py-4 text-2xl font-semibold">
                      {
                      typeof item[key] === 'number' ?
                      item[key].toLocaleString('en-US', { style: 'currency', currency: 'THB' }) :
                      item[key]
                      }
                    </td>
                  ))
                }
              </tr>
            ))
          }
          {
            total !== undefined && (
              <tr className="bg-pf-primary-2 text-pf-gray-100 text-center">
                <td className="px-6 py-4 text-2xl font-semibold">Total</td>
                {emptyCells}
                <td className="px-6 py-4 text-2xl font-semibold">
                  {total.toLocaleString('en-US', { style: 'currency', currency: 'THB' })}
                </td>
              </tr>
              )
          }
        </tbody>
      </table>
    </>
  );
}
