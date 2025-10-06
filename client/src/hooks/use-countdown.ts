import { useState, useEffect } from "react";

/**
 * Hook: useCountdownWithPercentage
 * @param {string} isoString - ISO 8601 string of target date/time
 * @param {number} startMinutes - initial countdown duration in minutes
 * @returns {{timeLeft: string, percentage: number}} countdown string and percentage
 */
export function useCountdown(isoString: string, startMinutes = 5) {
    const [timeLeft, setTimeLeft] = useState("00:00");
    const [percentage, setPercentage] = useState(0);

    useEffect(() => {
        const targetTime = new Date(isoString).getTime();
        const totalMs = startMinutes * 60 * 1000;

        function updateCountdown() {
            const now = Date.now();
            const diff = targetTime - now;

            if (diff <= 0) {
                setTimeLeft("00:00");
                setPercentage(0);
                return;
            }

            const minutes = Math.floor(diff / 1000 / 60);
            const seconds = Math.floor((diff / 1000) % 60);

            setTimeLeft(
                `${minutes.toString().padStart(2, "0")}:${seconds
                    .toString()
                    .padStart(2, "0")}`
            );

            const percent = (diff / totalMs) * 100;
            setPercentage(percent > 100 ? 100 : Math.round(percent));
        }

        updateCountdown(); // initial call
        const interval = setInterval(updateCountdown, 1000);

        return () => clearInterval(interval);
    }, [isoString, startMinutes]);

    return { timeLeft, percentage };
}
