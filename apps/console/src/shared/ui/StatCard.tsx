import { LucideIcon } from 'lucide-react';

interface StatCardProps {
  label: string;
  value: string;
  subText?: string;
  icon: LucideIcon;
  iconColor: string;
  iconBg: string;
}

export const StatCard = ({ label, value, subText, icon: Icon, iconColor, iconBg }: StatCardProps) => {
  return (
    <div className="bg-white p-5 rounded-[20px] shadow-sm flex items-center gap-5 transition-transform hover:scale-[1.02]">
      <div className={`w-14 h-14 rounded-full flex items-center justify-center ${iconBg} ${iconColor}`}>
        <Icon size={28} />
      </div>
      <div>
        <p className="text-sm text-[#A3AED0] font-medium mb-1">{label}</p>
        <h4 className="text-3xl font-bold text-[#2B3674] tracking-tight">{value}</h4>
        {subText && (
          <p className="text-xs text-[#A3AED0] mt-1 font-medium">
            {subText}
          </p>
        )}
      </div>
    </div>
  );
};
