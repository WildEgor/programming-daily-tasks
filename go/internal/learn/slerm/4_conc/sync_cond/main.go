package main

// TODO: Implement the task

/**
Изучите работу sync.Cond и на его основе сделайте реализацию уведомления N “спящих” горутин.
Пример: у нас есть worker pool, где каждая горутина постоянно опрашивает внешнюю систему (одну и ту же), берет оттуда задачи и их выполняет.
Мы знаем, что постоянный опрос это плохо, внешняя система упала.
У нас будет одна горутина, которая проверят liveness внешней системы и когда внешняя система оживет, мы должны уведомить все горутины, что они могут продолжать получать задачи из этой системы.
*/

func main() {

}
