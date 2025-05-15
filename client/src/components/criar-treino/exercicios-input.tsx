'use client'


interface Props {
    
}

const ExercicioInput: React.FC<Props> = ({  }) => {
    return (
        <div>
            <input type="text" name="exercicio" id="exercicio" placeholder="exercicio"/>
            <input type="number" name="series" id="series" min={1} defaultValue={1} />
        </div>
    );
};

export default ExercicioInput;