package models

type Muscle string
type ExerciseType string
type Equipment string
type Difficulty string

const (
	// Muscle
	MuscleAbductors       Muscle = "Абдукторы"
	MuscleAdductors       Muscle = "Аддукторы"
	MuscleBiceps          Muscle = "Бицепс"
	MuscleCalves          Muscle = "Икры"
	MuscleChest           Muscle = "Грудь"
	MuscleForearms        Muscle = "Предплечья"
	MuscleGlutes          Muscle = "Ягодицы"
	MuscleHips            Muscle = "Бедра"
	MuscleLatissimusDorsi Muscle = "Широчайшие мышцы спины"
	MuscleLowerBack       Muscle = "Нижняя часть спины"
	MuscleMiddleBack      Muscle = "Средняя часть спины"
	MuscleNeck            Muscle = "Шея"
	MusclePress           Muscle = "Пресс"
	MuscleQuadriceps      Muscle = "Квадрицепсы"
	MuscleShoulders       Muscle = "Плечи"
	MuscleTrapezius       Muscle = "Трапеции"
	MuscleTriceps         Muscle = "Трицепсы"

	// ExerciseType
	ExerciseTypeBasic      ExerciseType = "Базовое"
	ExerciseTypeIsolating  ExerciseType = "Изолирующее"

	// Equipment
	EquipmentBarbell      Equipment = "Штанга"
	EquipmentCable        Equipment = "Тросовые тренажеры"
	EquipmentDumbbells    Equipment = "Гантели"
	EquipmentFitball      Equipment = "Фитбол"
	EquipmentKettlebell   Equipment = "Гири"
	EquipmentMachine      Equipment = "Тренажер"
	EquipmentNone         Equipment = "Отсутствует"
	EquipmentOther        Equipment = "Другое"
	EquipmentOtherFitball Equipment = "Другое,Фитбол"
	EquipmentPullUpBar    Equipment = "Турник"
	EquipmentRackBarbell  Equipment = "Силовая рама,Штанга"
	EquipmentResistance   Equipment = "Эспандер"
	EquipmentSmithMachine Equipment = "Машина Смита"

	// Difficulty
	DifficultyBeginner     Difficulty = "Начинающий"
	DifficultyIntermediate Difficulty = "Средний"
	DifficultyAdvanced     Difficulty = "Профессионал"
)
