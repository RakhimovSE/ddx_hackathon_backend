package models

type Muscle string
type ExerciseType string
type Equipment string
type Difficulty string

const (
	// Muscle
	MusclePress           Muscle = "Пресс"
	MuscleLowerBack       Muscle = "Нижняя часть спины"
	MuscleTriceps         Muscle = "Трицепсы"
	MuscleShoulders       Muscle = "Плечи"
	MuscleQuadriceps      Muscle = "Квадрицепсы"
	MuscleBiceps          Muscle = "Бицепс"
	MuscleLatissimusDorsi Muscle = "Широчайшие мышцы спины"
	MuscleGlutes          Muscle = "Ягодицы"
	MuscleAdductors       Muscle = "Аддукторы"

	// ExerciseType
	ExerciseTypeBasic      ExerciseType = "Базовое"
	ExerciseTypeIsolating  ExerciseType = "Изолирующее"

	// Equipment
	EquipmentNone         Equipment = "Отсутствует"
	EquipmentKettlebell   Equipment = "Гири"
	EquipmentBarbell      Equipment = "Штанга"
	EquipmentDumbbells    Equipment = "Гантели"
	EquipmentMachine      Equipment = "Тренажер"
	EquipmentSmithMachine Equipment = "Машина Смита"
	EquipmentResistance   Equipment = "Эспандер"
	EquipmentOther        Equipment = "Другое"
	EquipmentFitball      Equipment = "Фитбол"
	EquipmentPullUpBar    Equipment = "Турник"

	// Difficulty
	DifficultyBeginner   Difficulty = "Начинающий"
	DifficultyIntermediate Difficulty = "Средний"
	DifficultyAdvanced   Difficulty = "Профессионал"
)
