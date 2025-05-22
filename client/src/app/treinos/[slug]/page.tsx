import Header from "@/components/header"
import TreinoSelecionado from "@/components/treinos/treino";

const Treino = async (props: { params: Promise<{ slug: string }> }) => {
    const params = await props.params;
    return (
        <>
            <Header />
            <TreinoSelecionado id={params.slug} />
        </>
    )
}

export default Treino