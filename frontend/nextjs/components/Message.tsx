export const Message = (props: any) => {
  if (props.type === 'error') {
    return (
      <div className="p-3 bg-red-200 border border-red-700 rounded text-red-700">
        {props.message}
      </div>
    );
  } else if (props.type === 'success') {
    return (
      <div className="p-3 bg-green-200 border border-green-700 rounded text-green-700">
        {props.message}
      </div>
    );
  } 

  return null;
}
