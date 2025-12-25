export const BillingWidget = () => {
  return (
    <div className="bg-white p-8 rounded-[30px] shadow-sm h-full flex flex-col justify-center">
      <div className="flex justify-between items-center mb-8">
        <h4 className="text-2xl font-bold text-[#2B3674]">공지사항</h4>
        <span className="text-sm font-bold text-[#4318FF] bg-[#F4F7FE] px-4 py-2 rounded-full">총 53건</span>
      </div>

      <div className="mb-10 text-center">
        <p className="text-xl text-[#A3AED0] mb-2 font-medium">총 청구액</p>
        <p className="text-4xl font-black text-[#2B3674] tracking-tight font-mono">₩37,499,242</p>
      </div>

      <div className="grid grid-cols-2 gap-5">
        <div className="bg-[#E6F9F3] p-5 rounded-2xl transition-transform hover:scale-105 cursor-default">
          <p className="text-sm font-bold text-[#05CD99] mb-2">완납</p>
          <p className="text-2xl font-bold text-[#2B3674]">29건</p>
        </div>
        <div className="bg-[#FFF9E5] p-5 rounded-2xl transition-transform hover:scale-105 cursor-default">
          <p className="text-sm font-bold text-[#FFB547] mb-2">수납대기중</p>
          <p className="text-2xl font-bold text-[#2B3674]">7건</p>
        </div>
        <div className="bg-[#FEEFEF] p-5 rounded-2xl transition-transform hover:scale-105 cursor-default">
          <p className="text-sm font-bold text-[#EE5D50] mb-2">미납</p>
          <p className="text-xl font-bold text-[#2B3674]">1건</p>
        </div>
        <div className="bg-[#E8F0FF] p-5 rounded-2xl transition-transform hover:scale-105 cursor-default">
          <p className="text-sm font-bold text-[#4318FF] mb-2">생성</p>
          <p className="text-xl font-bold text-[#2B3674]">16건</p>
        </div>
      </div>
    </div>
  );
};
