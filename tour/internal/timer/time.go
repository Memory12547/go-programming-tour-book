/*
 * @Author: Matt Meng
 * @Date: 2021-04-11 16:50:38
 * @LastEditors: Matt Meng
 * @LastEditTime: 2021-08-28 16:34:09
 * @Description: file content
 */
package timer

import "time"

func GetNowTime() time.Time {
    location, _ := time.LoadLocation("Asia/Shanghai")
    return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
    duration, err := time.ParseDuration(d)
    if err != nil {
        return time.Time{}, err
    }

    return currentTimer.Add(duration), nil
}
